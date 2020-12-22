
#include "suggest_service.hpp"
#include "utils.h"

#include <grpcpp/health_check_service_interface.h>
#include <grpcpp/ext/proto_server_reflection_plugin.h>

#include <algorithm>
#include <chrono>
#include <cstdint>
#include <fstream>
#include <memory>
#include <iostream>
#include <string>

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
using suggest::SuggestRequest;
using suggest::SuggestResponse;
using suggest::Suggest;
using suggest::Suggestion;
using nlohmann::json;

SuggestServiceImpl::SuggestServiceImpl() {
  std::thread([this]() -> void {
    std::ifstream file;
    while (true) {
      file.open("suggestions.json");
      std::unique_lock<std::shared_mutex> lock(collectionMutex_);
      file >> collection_;
      lock.unlock();
      std::cout << "sleeping for 15 min" << std::endl;
      file.close();
      std::this_thread::sleep_for(std::chrono::minutes(15));
    }
  }).detach();
}

Status SuggestServiceImpl::Input(ServerContext* context,
                                 const SuggestRequest* request,
                                 SuggestResponse* response) {
  google::protobuf::RepeatedPtrField<suggest::Suggestion> suggestions;
  std::shared_lock<std::shared_mutex> lock(collectionMutex_);
  suggestions.Reserve(collection_.size());
  for (const auto& record : collection_) {
    if (record.at("id").get<std::string>() == request->input()) {
      auto suggestion = suggestions.Add();
      *suggestion = record.get<suggest::Suggestion>();
    }
  }
  lock.unlock();
  std::sort(suggestions.begin(), suggestions.end(),
            [](const Suggestion& a, const Suggestion& b) -> bool {
              return a.position() < b.position();
            });
  uint32_t position = 1, cost = 0;
  using SizeType =
  typename google::protobuf::RepeatedPtrField<Suggestion>::size_type;
  for (SizeType i = 0; i != suggestions.size(); ++i) {
    if (cost != suggestions[i].position()) {
      suggestions[i].set_position(position++);
      cost = suggestions[i].position();
    } else {
      suggestions[i].set_position(position);
    }
  }
  *response->mutable_suggestions() = suggestions;
  return Status::OK;
}

void SuggestServer::Run(const std::string& server_address) {
  grpc::EnableDefaultHealthCheckService(true);
  grpc::reflection::InitProtoReflectionServerBuilderPlugin();
  ServerBuilder builder;
  // Listen on the given address without any authentication mechanism.
  builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
  // Register "service" as the instance through which we'll communicate with
  // clients. In this case it corresponds to an *synchronous* service.
  builder.RegisterService(&service);
  // Finally assemble the server.
  std::unique_ptr<Server> server(builder.BuildAndStart());
  std::cout << "Server listening on " << server_address << std::endl;

  // Wait for the server to shutdown. Note that some other thread must be
  // responsible for shutting down the server for this call to ever return.
  server->Wait();
}


#include "suggest_service.hpp"
#include "nlohmann/json.hpp"

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

namespace suggest {
  void from_json(const nlohmann::json& json_file, Suggestion& Sugg) {
    Sugg.set_position(json_file.at("cost").get<uint32_t>());
    Sugg.set_text(json_file.at("name").get<std::string>());
  }
}

SuggestServiceImpl::SuggestServiceImpl() {
  std::thread sync_thread([this]() {
    std::ifstream file;
    while (true) {
      file.open("suggestions.json");
      std::unique_lock<std::shared_mutex> lock(mutex);
      file >> collection;
      std::cout << "updated" << std::endl;
      lock.unlock();
      file.close();
      std::this_thread::sleep_for(std::chrono::minutes(15));
    }
  });
  sync_thread.detach();
}

Status SuggestServiceImpl::Input(ServerContext* context,
                                 const SuggestRequest* request,
                                 SuggestResponse* response) {
  google::protobuf::RepeatedPtrField<suggest::Suggestion> suggestions;
  std::shared_lock<std::shared_mutex> lock(mutex);
  suggestions.Reserve(collection.size());
  for (const auto& iter : collection) {
    if (iter.at("id").get<std::string>() == request->input()) {
      suggestions.Add(iter.get<suggest::Suggestion>());
    }
  }
  lock.unlock();
  std::sort(suggestions.begin(), suggestions.end(),
            [](const Suggestion& a, const Suggestion& b) {
              return a.position() < b.position();
            });
  uint32_t position = 0, cost = 0;
  for (auto& iter : suggestions) {
    if (cost != iter.position()) {
      position++;
      iter.set_position(position);
      cost = iter.position();
    } else {
      iter.set_position(position);
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

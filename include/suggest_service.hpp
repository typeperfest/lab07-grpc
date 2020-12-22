// Copyright 2020 Your Name <your_email>

#ifndef SOURCES_SUGGEST_SERVICE_HPP_
#define SOURCES_SUGGEST_SERVICE_HPP_

#include <grpcpp/grpcpp.h>

#include <suggest.grpc.pb.h>

#include <nlohmann/json.hpp>

#include <chrono>
#include <mutex>
#include <thread>
#include <shared_mutex>
#include <string>

class SuggestServiceImpl final : public suggest::Suggest::Service {
 public:
  SuggestServiceImpl();
 private:
  grpc::Status Input(grpc::ServerContext* context,
                     const suggest::SuggestRequest* request,
                     suggest::SuggestResponse* response) override;
  std::shared_mutex collectionMutex_;
  nlohmann::json collection_;
}; // class SuggestServiceImpl

class SuggestServer final {
 public:
  inline SuggestServer() noexcept = default;
  void Run(const std::string& server_address);
 private:
  SuggestServiceImpl service;
}; // class SuggestService

#endif // SOURCES_SUGGEST_SERVICE_HPP_
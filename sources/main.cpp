#include "suggest_service.hpp"

int main(int argc, char** argv) {
  SuggestServer server;
  server.Run("0.0.0.0:9090");
  return 0;
}
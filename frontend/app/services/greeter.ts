import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { GreeterService } from "../gen/greeter/v1/greeter_pb";

// 環境変数からバックエンドURLを取得
const getBackendUrl = () => {
  const backendUrl = import.meta.env.VITE_BACKEND_URL;
  
  // 環境変数が設定されていない場合は相対パス（本番環境用）
  if (!backendUrl) {
    return "/api";
  }
  
  return backendUrl;
};

// トランスポートの作成
const transport = createConnectTransport({
  baseUrl: getBackendUrl(),
});

export const greeterClient = createClient(GreeterService, transport);

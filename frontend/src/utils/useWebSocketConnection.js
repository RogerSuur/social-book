import { useEffect } from "react";
import useWebSocket from "react-use-websocket";

export const useWebSocketConnection = (socketUrl, onMessage) => {
  const { sendJsonMessage, lastJsonMessage } = useWebSocket(socketUrl, {
    onOpen: () => console.log("WebSocket opened"),
    share: true,
  });

  useEffect(() => {
    sendJsonMessage("hello");
  }, [sendJsonMessage]);

  useEffect(() => {
    if (onMessage && typeof onMessage === "function") {
      onMessage(lastJsonMessage);
    }
  }, [lastJsonMessage, onMessage]);
};

import { useEffect } from "react";
import useWebSocket from "react-use-websocket";

export const useWebSocketConnection = (socketUrl, message) => {
  const { sendJsonMessage, lastJsonMessage } = useWebSocket(socketUrl, {
    onOpen: () => console.log("WebSocket opened"),
    share: true,
  });

  useEffect(() => {
    sendJsonMessage(message);
  }, [sendJsonMessage]);

  // useEffect(() => {
  //   if (onMessage && typeof onMessage === "function") {
  //     onMessage(lastJsonMessage);
  //   }
  // }, [lastJsonMessage, onMessage]);
};

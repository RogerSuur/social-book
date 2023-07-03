import useWebSocket from "react-use-websocket";

const useWebSocketConnection = (socketUrl) => {
  const { sendJsonMessage, lastJsonMessage } = useWebSocket(socketUrl, {
    onOpen: () => console.log("WebSocket opened"),
    share: true,
  });

  // useEffect(() => {
  //   sendJsonMessage(message);
  // }, [sendJsonMessage]);

  // useEffect(() => {
  //   if (onMessage && typeof onMessage === "function") {
  //     onMessage(lastJsonMessage);
  //   }
  // }, [lastJsonMessage, onMessage]);

  return { sendJsonMessage, lastJsonMessage };
};

export default useWebSocketConnection;

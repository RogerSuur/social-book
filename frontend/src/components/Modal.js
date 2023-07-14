import Popup from "reactjs-popup";
import "../style.css";

const Modal = ({ text, children, open, onClose }) => (
  <Popup
    contentStyle={{
      display: "flex",
      flexWrap: "wrap",
      justifyContent: "center",
      overflow: "auto",
      backgroundColor: "white",
      width: "20%",
      transform: "translate(-40%, -10%)",
    }}
    open={open}
    onClose={onClose}
    modal
  >
    <span> {children}</span>
  </Popup>
);

export default Modal;

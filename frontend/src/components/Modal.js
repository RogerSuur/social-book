import Popup from "reactjs-popup";
import "../style.css";

const Modal = ({ text, children, open, onClose }) => (
  <Popup
    contentStyle={{
      
  justifyContent: "center",
  alignItems: "center",
  backgroundColor: "lightgray",
  marginTop: "15vh",
  overflow: "auto",
  maxHeight: "60vh",
  borderRadius: "10px",
  boxShadow: "0 4px 8px rgba(0, 0, 0, 0.1)",
    }}
    open={open}
    onClose={onClose}
    modal
  >
    <span> {children}</span>
  </Popup>
);

export default Modal;

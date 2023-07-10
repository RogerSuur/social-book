import Popup from "reactjs-popup";
import "reactjs-popup/dist/index.css";

const Modal = ({ text, children, open, onClose }) => (
  <Popup
    style={{ height: "300px", overflow: "initial" }}
    open={open}
    onClose={onClose}
    modal
  >
    <span style={{ overflow: "auto" }}> {children}</span>
  </Popup>
);

export default Modal;

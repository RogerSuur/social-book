import Popup from "reactjs-popup";
import "reactjs-popup/dist/index.css";

const Modal = ({ text, children, open, onClose }) => (
  <Popup open={open} onClose={onClose} modal>
    <span> {children}</span>
  </Popup>
);

export default Modal;

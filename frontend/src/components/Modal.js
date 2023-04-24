import Popup from "reactjs-popup";
import "reactjs-popup/dist/index.css";

const Modal = ({ text, children }) => (
  <Popup trigger={<button> {text} </button>} modal>
    <span> {children}</span>
  </Popup>
);

export default Modal;

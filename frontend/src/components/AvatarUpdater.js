import AvatarEditor from "react-avatar-editor";
import { useState, useRef } from "react";
import FileUploader from "./FileUploader";
import axios from "axios";

const AvatarUpdater = ({ url, onUploadSuccess }) => {
  const editorRef = useRef();
  const [selectedImage, setSelectedImage] = useState(null);
  const [errMsg, setErrMsg] = useState("");

  const handleClick = async () => {
    const canvas = editorRef.current.getImage();

    canvas.toBlob(async (blob) => {
      const formData = new FormData();
      formData.append("image", blob, "avatar.jpg");

      try {
        await axios.post(url, formData, { withCredentials: true });
        onUploadSuccess();
        console.log("Image uploaded successfully!");
      } catch (err) {
        if (!err?.response) {
          setErrMsg("No Server Response");
        } else if (err.response?.status > 200) {
          //handle image size errors
          setErrMsg("Internal Server Error");
        }
      }
    }, "image/jpeg");
  };

  return (
    <>
      <AvatarEditor
        ref={editorRef}
        image={
          selectedImage
            ? selectedImage
            : `${process.env.PUBLIC_URL}/defaultuser.jpg`
        }
        width={250}
        height={250}
        color={[255, 255, 255, 0.6]} // RGBA
        scale={1.2}
        rotate={0}
      />
      <button onClick={handleClick}>Save image</button>
      {errMsg && <h3>{errMsg}</h3>}
      <FileUploader
        onFileSelectSuccess={(file) => {
          setSelectedImage(file);
          setErrMsg("");
        }}
        onFileSelectError={({ error }) => setErrMsg(error)}
      />
    </>
  );
};

export default AvatarUpdater;

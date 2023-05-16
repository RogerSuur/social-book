import AvatarEditor from "react-avatar-editor";
import { useState, useRef } from "react";
import FileUploader from "./FileUploader";
import axios from "axios";

const IMAGE_UPLOAD_URL = "http://localhost:8000/profile/update/avatar";

const AvatarUpdater = ({ onUploadSuccess }) => {
  const editorRef = useRef();
  const [selectedImage, setSelectedImage] = useState(null);
  const [errMsg, setErrMsg] = useState("");

  const handleClick = async () => {
    const canvas = editorRef.current.getImage();

    canvas.toBlob(async (blob) => {
      // Create a FormData object and append the blob as a file
      const formData = new FormData();
      formData.append("image", blob, "avatar.png");
      try {
        // Send the image data to the server using Axios
        await axios.post(IMAGE_UPLOAD_URL, formData, { withCredentials: true });
        onUploadSuccess();
        console.log("Image uploaded successfully!");
      } catch (err) {
        if (!err?.response) {
          setErrMsg("No Server Response");
        } else if (err.response?.status > 200) {
          setErrMsg("Internal Server Error");
        }
      }
    });
  };

  return (
    <>
      <AvatarEditor
        ref={editorRef}
        image={
          selectedImage
            ? selectedImage
            : "https://hopatcongpolice.org/wp-content/uploads/2019/03/default-person.png"
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
        onFileSelectSuccess={(file) => setSelectedImage(file)}
        onFileSelectError={({ error }) => setErrMsg(error)}
      />
    </>
  );
};

export default AvatarUpdater;

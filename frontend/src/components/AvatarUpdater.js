import AvatarEditor from "react-avatar-editor";
import { useState, useRef } from "react";
import FileUploader from "./FileUploader";

const AvatarUpdater = ({ children }) => {
  const editorRef = useRef();
  const [selectedImage, setSelectedImage] = useState(null);
  const [errMsg, setErrMsg] = useState("");
  const [croppedImage, setCroppedImage] = useState(null);

  const handleClick = () => {
    const canvas = editorRef.current.getImage().toDataURL();
    setCroppedImage(canvas);
  };

  console.log(croppedImage);

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
        borderRadius={50}
        color={[255, 255, 255, 0.6]} // RGBA
        scale={1.2}
        rotate={0}
      />
      {croppedImage && (
        <img style={{ width: "20vw", height: "auto" }} src={croppedImage} />
      )}
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

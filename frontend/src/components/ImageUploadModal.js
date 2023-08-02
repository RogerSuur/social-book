import React, { useState } from "react";
import Popup from "reactjs-popup";
import "../style.css";

const ImageUploadModal = ({ open, onClose, onImageUpload }) => {
  const [selectedImage, setSelectedImage] = useState(null);

  const handleFileInput = (event) => {
    const file = event.target.files[0];
    setSelectedImage(file);
  };

  const handleUpload = () => {
    onImageUpload(selectedImage);
    setSelectedImage(null);
    onClose();
  };

  return (
    <Popup
      contentStyle={{
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "rgb(224, 224, 224)",
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
      <div className="image-upload-modal">
        <h2>Upload Image</h2>
        <input type="file" accept="image/*" onChange={handleFileInput} />
        {selectedImage && (
          <div>
            <img
              src={URL.createObjectURL(selectedImage)}
              alt="Selected"
              className="profile-image"
            />
            <button onClick={handleUpload}>Upload</button>
          </div>
        )}
        <button onClick={onClose}>Cancel</button>
      </div>
    </Popup>
  );
};

export default ImageUploadModal;

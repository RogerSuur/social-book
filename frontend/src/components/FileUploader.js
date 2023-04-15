import React, { useRef } from "react";

const FileUploader = ({ onFileSelectSuccess, onFileSelectError }) => {
  const fileInput = useRef(null);

  const handleFileInput = (event) => {
    const file = event.target.files[0];
    if (file.size > 10240)
      onFileSelectError({ error: "File size cannot exceed more than 10MB" });
    else onFileSelectSuccess(file);
  };

  return (
    <div className="file-uploader">
      <input type="file" onChange={handleFileInput} />
      {/* <button
        onClick={(e) => fileInput.current && fileInput.current.click()}
        className="btn-primary"
      /> */}
    </div>
  );
};

export default FileUploader;

import React from "react";

const FileUploader = ({ onFileSelectSuccess, onFileSelectError }) => {
  const handleFileInput = (event) => {
    const file = event.target.files[0];
    if (file.size > 5242880)
      onFileSelectError({ error: "File size cannot exceed 5MB" });
    else onFileSelectSuccess(file);
  };

  return (
    <div className="file-uploader">
      <input type="file" onChange={handleFileInput} />
    </div>
  );
};

export default FileUploader;

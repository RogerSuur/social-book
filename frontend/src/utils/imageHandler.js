const ImageSource = (path, defaultImage) =>
  path
    ? `${process.env.PUBLIC_URL}/images/${path}`
    : `${process.env.PUBLIC_URL}/${defaultImage}`;

const ImageHandler = (path, defaultImage, className) => {
  return <img className={className} src={ImageSource(path, defaultImage)} />;
};

export default ImageHandler;

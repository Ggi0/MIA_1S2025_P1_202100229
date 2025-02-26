// subir un archivo desde su computadora.

export function FileUpload({ onFileUpload }) {
  
    const handleFileChange = (event) => {
    const file = event.target.files[0];

    if (file) {
      onFileUpload(file);
    }
  };

  return (
    <div>
      <input type="file" onChange={handleFileChange} />
    </div>
  );
}

export default FileUpload;

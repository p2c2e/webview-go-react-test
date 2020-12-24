import React, { useRef } from 'react';
import { FileDrop } from 'react-file-drop';
import './App.css';
//import $ from 'jquery';

export const App = () => {
  const styles = { border: '1px solid black', width: 600, color: 'black', padding: 20 };
  const fileInputRef = useRef(null);
  const submitButtonRef = useRef(null);
  const formRef = useRef( null )
  const onFileInputChange = (event) => {
    const { files } = event.target;
    // do something with your files...
    for (var i = 0; i < files.length; i++) {
        console.log("--------> "+files[i]);
    }
  }
  const onFileInputChange2 = () => {
    // do something with your files...
    console.log(fileInputRef.current.files);
    console.log("submitting");
    // submitButtonRef.current.click();
    formRef.current.submit(); // dispatchEvent(new Event('submit', { cancelable : true }))
    console.log("submitted");
  }
  const onTargetClick = () => {
    fileInputRef.current.click()
  }
  return (
    <div>
      <h1>React File Drop demo</h1>
      <div style={styles}>
        <FileDrop
          onDrop={(files, event) => { console.log('onDrop!', files, event); fileInputRef.current.files=files; onFileInputChange2(); } }
          onTargetClick={onTargetClick}
        >
          Drop some files here!
        </FileDrop>
      </div>
      <form id="fileselect" // onSubmit={ (event) => event.preventDefault() }
        action="/recvfile" method="post" encType="multipart/form-data"
      ref={formRef}>
      <input
        onChange={onFileInputChange}
        ref={fileInputRef}
        type="file"
        className="hidden"
        id="fileinput"
      />
      <input type="submit" ref={submitButtonRef} value="submit"/>
      </form>

      <form encType="multipart/form-data" action="/upload" method="post">
        <input type="file" name="uploadfile"/>
        <input type="hidden" name="token" value="{{.}}"/>
        <input type="submit" value="upload"/>
      </form>
    </div>
  );
};

export default App;

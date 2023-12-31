import React, { useState } from "react";
function Modal({children, isShow, setIsShow, title}) {
  const onCloseClick = () => {
    setIsShow(false)
  }

  return(
    <div id="myModal" className="modal" style={{display: isShow ? "block" : "none"}}>

    <div className="modal-content">
      <span className="close" onClick={onCloseClick}>&times;</span>
      {title != null ? <div className="modal-title">{title}</div> : ''}
      {children}
    </div>
  
  </div>
  );
}
export default Modal;
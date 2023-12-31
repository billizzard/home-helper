function Title({title}) {
  return <div className="title-box">
    <a className="title-btn shadow title-btn fix-blink" href='/'>
        <i className="bi bi-house-door"></i>
    </a>
    <h1 className="page-title">{title}</h1>
  </div>
  ;
}
export default Title;
function Alert({alert}) {
  return <div className={"alerts " + alert.type + ' ' + (alert.message != null ? 'show' : 'hide')}>
        {alert.message}
      </div>
  ;
}
export default Alert;
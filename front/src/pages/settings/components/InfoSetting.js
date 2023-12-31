function InfoSetting({title, value, note = ""}) {
    return <div className="st-option">
        <div className="st-option_title">
            <label>{title}</label>
            <div>{value}</div>
        </div>
        <div className="st-note">{note}</div>
    </div>
}

export default InfoSetting
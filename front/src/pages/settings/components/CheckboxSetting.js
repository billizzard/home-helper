import { useState } from "react"

function CheckboxSetting({title, note, startValue, onChange}) {
    const [value, setValue] = useState(startValue)

    const onChangeOption = () => {
        setValue(!value)
        if (onChange) {
            onChange(!value)
        }
    }

    return <div className="st-option">
        <div className="st-option_title">
            <label>{title}</label>
            <input type="checkbox" checked={value} onChange={onChangeOption}/> 
        </div>
        <div className="st-note">{note}</div>
    </div>
}

export default CheckboxSetting
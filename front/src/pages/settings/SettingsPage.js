import { useContext } from "react";
import { AppContext } from "../../context/appContext";
import Title from "../layout/Title";
import CheckboxSetting from "./components/CheckboxSetting";
import InfoSetting from "./components/InfoSetting";

function SettingsPage() {
  const app = useContext(AppContext);

  const onPreviewChange = (value) => {
    localStorage.setItem('imagePreview', value);
    const newData = app.appData
    newData.hasPreview = value
    app.setAppData(newData)
  }

  return (
      <>
          <Title title="Settings"/>
          <div className="st-body">
              <CheckboxSetting
                  startValue={app.appData.hasPreview}
                  title="Show image preview"
                  note="This may slow down a page with large images or a lot of images"
                  onChange={onPreviewChange}
              />
          </div>
          <div className="st-body">
              <InfoSetting title="Server address" value={window.appConfig.backendUrl} note="Enter the address on other devices to share files"/>
          </div>
      </>
  );
}

export default SettingsPage;

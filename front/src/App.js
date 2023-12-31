import { BrowserRouter, Route, Routes } from 'react-router-dom';
import MenuPage from './pages/mainMenu/MenuPage';
import FileListPage from './pages/fileList/FileListPage';
import Layout from './pages/layout/Layout';
import Error404Page from './pages/errors/Error404Page';
import SettingsPage from './pages/settings/SettingsPage';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout><MenuPage /></Layout>}/>
        <Route path="/files/*" element={<Layout><FileListPage /></Layout>} />
        <Route path="/settings" element={<Layout><SettingsPage /></Layout>} />
        <Route path="/404" element={<Layout><Error404Page /></Layout>} />
        <Route path="*" element={<Layout><Error404Page /></Layout>} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

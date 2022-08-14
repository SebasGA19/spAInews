import { BrowserRouter, Routes, Route } from 'react-router-dom';

import { Layout } from './pages/layout';
import { Home } from './pages/home';
import { Login } from './pages/login';
import { Register } from './pages/register';
import { ConfirmRegister } from './pages/confirm-register';
import { Settings} from './pages/settings';
import { ChangePassword} from './pages/change-password';

function App(){
  return(
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="login" element={<Login />} />
          <Route path="register" element={<Register />} />
          <Route path="confirm-register" element={<ConfirmRegister />} />
          <Route path="settings" element={<Settings />} />
          <Route path="change-password" element={<ChangePassword />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App;

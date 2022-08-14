import { BrowserRouter, Routes, Route } from 'react-router-dom';

import { Layout } from './pages/layout';
import { Home } from './pages/home';
import { Login } from './pages/login';
import { Register } from './pages/register';
import { ConfirmRegister } from './pages/confirm-register';

function App(){
  return(
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="login" element={<Login />} />
          <Route path="register" element={<Register />} />
          <Route path="confirm-register" element={<ConfirmRegister />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App;

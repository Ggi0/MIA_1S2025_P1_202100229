import Principal from "./pages/principal";
import Segundaria from "./pages/segundaria"
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { Navigation } from "./components/Navigation";

function App() {
  return (
    <BrowserRouter>
      <Navigation/>
      <Routes>
        <Route path="/" element={<Principal />} />
        <Route path="/info" element={<Segundaria />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

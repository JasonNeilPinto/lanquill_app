import { React } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import "./App.css";
// Swiper CSS
import "swiper/css";
import "swiper/css/navigation";

//bootstrap
import "bootstrap/dist/js/bootstrap.bundle";
import ScrollToTop from "./components/common/ScrollToTop";
import HomePage from "./components/themes/HomePage";
import About from "./components/pages/services/About";

function App() {
  return (
    <>
      <Router>
        <ScrollToTop />
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/about-us" element={<About />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;

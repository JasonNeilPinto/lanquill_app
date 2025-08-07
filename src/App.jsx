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
import DataAnalytics from "./components/services/DataAnaytics/DataAnalyticsService";
import CyberSecurityService from "./components/services/CyberSecurity/CyberSecurityService";
import CloudServices from "./components/services/CloudServices/CloudService";
import AppliedAIServices from "./components/services/AppliedAI/AppliedAIService";
import SecurityTestingService from "./components/services/SecurityTesting/SecurityTestingService";
import ManagedServices from "./components/services/ManagedServices/ManagedServices";
import Contact from "./components/pages/services/Contact";

function App() {
  return (
    <>
      <Router>
        <ScrollToTop />
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/about-us" element={<About />} />

          <Route path="/services/data-analytics" element={<DataAnalytics />} />
          <Route
            path="/services/cyber-security"
            element={<CyberSecurityService />}
          />
          <Route path="/services/cloud-services" element={<CloudServices />} />
          <Route path="/services/applied-ai" element={<AppliedAIServices />} />
          <Route
            path="/services/security-testing"
            element={<SecurityTestingService />}
          />
          <Route
            path="/services/managed-services"
            element={<ManagedServices />}
          />

          <Route path="/contact-us" element={<Contact />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;

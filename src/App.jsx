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
import Blogs from "./components/pages/blog/Blogs";
import SingleBlog from "./components/pages/blog/SingleBlog";
import SecondBlog from "./components/pages/blog/SecondBlog";
import ShapingIndiaBlog from "./components/pages/blog/ShapingIndiaBlog";
import TheWarForTalentBlog from "./components/pages/blog/TheWarForTalentBlog";
import GenerativeAiBlog from "./components/pages/blog/GenerativeAiBlog";
import EmergenceOfGenAiBlog from "./components/pages/blog/EmergenceOfGenAiBlog";
import GenAiRevolutionizesBlog from "./components/pages/blog/GenAiRevolutionizesBlog";
import EcommerceBlog from "./components/pages/blog/EcommerceBlog";
import RequestDemo from "./components/pages/services/RequestDemo";

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

          <Route path="/blogs" element={<Blogs />} />
          <Route path="/blog-single" element={<SingleBlog />} />
          <Route path="/blog-second" element={<SecondBlog />} />
          <Route path="/blog-ShapingIndia" element={<ShapingIndiaBlog />} />
          <Route
            path="/blog-theWarfortalent"
            element={<TheWarForTalentBlog />}
          />
          <Route path="/blog-generativeai" element={<GenerativeAiBlog />} />
          <Route
            path="/blog-emergenceofgenai"
            element={<EmergenceOfGenAiBlog />}
          />
          <Route
            path="/blog-genairevolutionizes"
            element={<GenAiRevolutionizesBlog />}
          />
          <Route path="/blog-ecommerce" element={<EcommerceBlog />} />

          <Route path="/request-demo" element={<RequestDemo />} />

          <Route path="/contact-us" element={<Contact />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;

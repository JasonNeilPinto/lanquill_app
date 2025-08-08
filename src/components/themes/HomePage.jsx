import React, { useEffect } from "react";
import Layout from "../../layout/Layout";
import PageMeta from "../common/PageMeta";
import Navbar from "../../layout/Header/Navbar";
// import HeroTen from "./HeroPage";
// import CustomerLogoSlider from "../customer/CustomerLogoSlider";
import CyberAbout from "../about/CyberAbout";
import CyberService from "../services/CyberService";
import CyberCta from "../cta/CyberCta";
import CyberVideoPromo from "../promo/CyberVideoPromo";
import TestimonialFour from "../testimonials/TestimonialFour";
import FooterTwo from "../../layout/Footer/FooterTwo";
import DesktopHome from "../../desktopApp/DesktopHome";
import Products from "../products/Products";
import LatestBlog from "../blogs/LatestBlog";

const HomePage = () => {
  useEffect(() => {
    const hash = window.location.hash;
    if (hash) {
      const scrollAndActivate = () => {
        const section = document.querySelector("#div-products");
        if (section) section.scrollIntoView({ behavior: "smooth" });

        const tabBtn = document.querySelector(`a[href="${hash}"]`);
        if (tabBtn instanceof HTMLElement) tabBtn.click();
      };

      // Small timeout to wait for DOM to load
      setTimeout(scrollAndActivate, 200);
    }
  }, []);

  return (
    <Layout>
      <PageMeta title="Netanalytiks" />
      <Navbar />
      <DesktopHome />
      <CyberAbout />
      <CyberService />
      <CyberCta />
      <Products />
      <CyberVideoPromo />
      {/* <PriceFour /> */}
      {/* <CyberStore /> */}
      <TestimonialFour />
      {/* <CyberFaq /> */}
      <LatestBlog />
      <FooterTwo />
    </Layout>
  );
};

export default HomePage;

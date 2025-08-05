import React from "react";
import Layout from "../../layout/Layout";
import PageMeta from "../common/PageMeta";
import Navbar from "../../layout/Header/Navbar";
// import HeroTen from "./HeroPage";
// import CustomerLogoSlider from "../customer/CustomerLogoSlider";
import CyberAbout from "../about/CyberAbout";
import CyberService from "../services/CyberService";
import CyberCta from "../cta/CyberCta";
import CyberVideoPromo from "../promo/CyberVideoPromo";
import PriceFour from "../prices/PriceFour";
import CyberStore from "../others/CyberStore";
import TestimonialFour from "../testimonials/TestimonialFour";
import CyberFaq from "../faqs/CyberFaq";
import CyberBlog from "../blogs/CyberBlog";
import FooterTwo from "../../layout/Footer/FooterTwo";
import DesktopHome from "../../desktopApp/DesktopHome";
import Products from "../products/Products";

const HomePage = () => {
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
      <PriceFour />
      <CyberStore />
      <TestimonialFour />
      <CyberFaq />
      <CyberBlog />
      <FooterTwo />
    </Layout>
  );
};

export default HomePage;

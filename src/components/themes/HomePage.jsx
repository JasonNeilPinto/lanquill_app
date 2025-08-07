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
import TestimonialFour from "../testimonials/TestimonialFour";
import FooterTwo from "../../layout/Footer/FooterTwo";
import DesktopHome from "../../desktopApp/DesktopHome";
import Products from "../products/Products";
import BlogItems from "../blogs/BlogItem";

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
      {/* <PriceFour /> */}
      {/* <CyberStore /> */}
      <TestimonialFour />
      {/* <CyberFaq /> */}
      <BlogItems />
      <FooterTwo />
    </Layout>
  );
};

export default HomePage;

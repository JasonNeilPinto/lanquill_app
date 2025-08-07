import React from "react";
import Layout from "../../../layout/Layout";
import Navbar from "../../../layout/Header/Navbar";
import Feature from "./Feature";
import TestimonialTwo from "../../testimonials/TestimonialTwo";
import CtaTwo from "../../cta/CtaTwo";
import FooterOne from "../../../layout/Footer/FooterOne";
import FeatureTop from "./FeatureTop";
import PageHeader from "../../common/PageHeader";
import BlogItems from "../../blogs/BlogItem";

const CyberSecurityService = () => {
  return (
    <Layout>
      <Navbar navDark />
      <PageHeader
        title="Cyber Security"
        desc="Protecting What Matters Most - Your Digital Assets"
      />
      <FeatureTop />
      <Feature />
      <BlogItems />
      <TestimonialTwo bgWhite />
      <CtaTwo />
      <FooterOne footerLight />
    </Layout>
  );
};

export default CyberSecurityService;

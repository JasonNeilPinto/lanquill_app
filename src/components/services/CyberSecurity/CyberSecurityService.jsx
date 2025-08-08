import React from "react";
import Layout from "../../../layout/Layout";
import Navbar from "../../../layout/Header/Navbar";
import Feature from "./Feature";
import TestimonialTwo from "../../testimonials/TestimonialTwo";
import CtaTwo from "../../cta/CtaTwo";
import FeatureTop from "./FeatureTop";
import PageHeader from "../../common/PageHeader";
import FooterTwo from "../../../layout/Footer/FooterTwo";
import LatestBlog from "../../blogs/LatestBlog";

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
      <LatestBlog />
      <TestimonialTwo bgWhite />
      <CtaTwo />
      <FooterTwo />
    </Layout>
  );
};

export default CyberSecurityService;

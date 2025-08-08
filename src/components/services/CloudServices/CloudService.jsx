import React from "react";
import CtaTwo from "../../cta/CtaTwo";
import Navbar from "../../../layout/Header/Navbar";
import Layout from "../../../layout/Layout";
import Feature from "./Feature";
import TestimonialTwo from "../../testimonials/TestimonialTwo";
import FeatureTop from "./FeatureTop";
import PageHeader from "../../common/PageHeader";
import FooterTwo from "../../../layout/Footer/FooterTwo";
import LatestBlog from "../../blogs/LatestBlog";

const CloudServices = () => {
  return (
    <Layout>
      <Navbar navDark />
      <PageHeader
        title="Cloud Services"
        desc="Drive Agility and Innovation with Cloud Services"
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

export default CloudServices;

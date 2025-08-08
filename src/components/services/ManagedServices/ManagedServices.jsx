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

const ManagedServices = () => {
  return (
    <Layout>
      <Navbar navDark />
      <PageHeader
        title="Managed Services"
        desc="Cost-efficient, scalable, and high-quality IT Managed Services."
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

export default ManagedServices;

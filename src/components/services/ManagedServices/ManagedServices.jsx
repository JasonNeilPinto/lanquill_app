import React from "react";
import CtaTwo from "../../cta/CtaTwo";
import FooterOne from "../../../layout/Footer/FooterOne";
import Navbar from "../../../layout/Header/Navbar";
import Layout from "../../../layout/Layout";
import Feature from "./Feature";
import TestimonialTwo from "../../testimonials/TestimonialTwo";
import FeatureTop from "./FeatureTop";
import PageHeader from "../../common/PageHeader";
import BlogItems from "../../blogs/BlogItem";

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
      <BlogItems />
      <TestimonialTwo bgWhite />
      <CtaTwo />
      <FooterOne footerLight />
    </Layout>
  );
};

export default ManagedServices;

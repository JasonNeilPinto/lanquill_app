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

const AppliedAIServices = () => {
  return (
    <Layout>
      <Navbar navDark />
      <PageHeader
        title="Applied AI "
        desc="Transforming modern enterprises innovate, automate, and grow"
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

export default AppliedAIServices;

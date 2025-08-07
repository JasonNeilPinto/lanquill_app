import React from "react";
import CtaTwo from "../../cta/CtaTwo";
import Navbar from "../../../layout/Header/Navbar";
import Layout from "../../../layout/Layout";
import Feature from "./Feature";
import TestimonialTwo from "../../testimonials/TestimonialTwo";
import FeatureTop from "./FeatureTop";
import PageHeader from "../../common/PageHeader";
import BlogItems from "../../blogs/BlogItem";
import FooterTwo from "../../../layout/Footer/FooterTwo";

const DataAnalytics = () => {
  return (
    <Layout>
      <Navbar navDark />
      <PageHeader
        title="Data & Analytics"
        desc="Build, Transform, Scale - Your Data, Our Expertise"
      />
      <FeatureTop />
      <Feature />
      <BlogItems />
      <TestimonialTwo bgWhite />
      <CtaTwo />
      <FooterTwo />
    </Layout>
  );
};

export default DataAnalytics;

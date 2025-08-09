import React from "react";
import CtaTwo from "../../cta/CtaTwo";
import Navbar from "../../../layout/Header/Navbar";
import Layout from "../../../layout/Layout";
import Feature from "./Feature";
import FeatureTop from "./FeatureTop";
import PageHeader from "../../common/PageHeader";
import FooterTwo from "../../../layout/Footer/FooterTwo";
import LatestBlog from "../../blogs/LatestBlog";

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
      <LatestBlog />
      <CtaTwo />
      <FooterTwo />
    </Layout>
  );
};

export default DataAnalytics;

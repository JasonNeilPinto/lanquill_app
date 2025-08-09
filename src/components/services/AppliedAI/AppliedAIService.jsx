import React from "react";
import CtaTwo from "../../cta/CtaTwo";
import Navbar from "../../../layout/Header/Navbar";
import Layout from "../../../layout/Layout";
import Feature from "./Feature";
import FeatureTop from "./FeatureTop";
import PageHeader from "../../common/PageHeader";
import FooterTwo from "../../../layout/Footer/FooterTwo";
import LatestBlog from "../../blogs/LatestBlog";

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
      <LatestBlog />
      <CtaTwo />
      <FooterTwo />
    </Layout>
  );
};

export default AppliedAIServices;

import React from "react";
import Layout from "../../../layout/Layout";
import PageMeta from "../../common/PageMeta";
import Navbar from "../../../layout/Header/Navbar";
import PageHeader from "../../common/PageHeader";
import LatestBlog from "../../blogs/LatestBlog";
import CtaTwo from "../../cta/CtaTwo";
import FooterOne from "../../../layout/Footer/FooterOne";
import GenAiRevolutionizes from "../../blogs/GenAiRevolutionizes";

const GenAiRevolutionizesBlog = () => {
  return (
    <Layout>
      <PageMeta title="Blog Details - Software &amp; IT Solutions HTML Template" />
      <Navbar navDark />
      <PageHeader title="How Gen AI Revolutionizes Coding: The Future of AI-Powered Coding Assistants" />
      <GenAiRevolutionizes />
      {/* <NewsLetter /> */}
      <LatestBlog />
      <CtaTwo />
      <FooterOne footerLight />
    </Layout>
  );
};

export default GenAiRevolutionizesBlog;

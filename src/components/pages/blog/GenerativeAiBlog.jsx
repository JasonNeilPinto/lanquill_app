import React from "react";
import Layout from "../../../layout/Layout";
import PageMeta from "../../common/PageMeta";
import Navbar from "../../../layout/Header/Navbar";
import PageHeader from "../../common/PageHeader";
import LatestBlog from "../../blogs/LatestBlog";
import CtaTwo from "../../cta/CtaTwo";
import GenerativeAi from "../../blogs/GenerativeAi";
import FooterTwo from "../../../layout/Footer/FooterTwo";

const GenerativeAiBlog = () => {
  return (
    <Layout>
      <PageMeta title="Blog Details - Software &amp; IT Solutions HTML Template" />
      <Navbar navDark />
      <PageHeader title="How Generative AI Will Transform the Legal Industry" />
      <GenerativeAi />
      {/* <NewsLetter /> */}
      <LatestBlog />
      <CtaTwo />
      <FooterTwo />
    </Layout>
  );
};

export default GenerativeAiBlog;

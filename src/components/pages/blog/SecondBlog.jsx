import React from "react";
import Layout from "../../../layout/Layout";
import PageMeta from "../../common/PageMeta";
import Navbar from "../../../layout/Header/Navbar";
import PageHeader from "../../common/PageHeader";
import BlogDetailsSecond from "../../blogs/BlogDetailsSecond";
import LatestBlog from "../../blogs/LatestBlog";
import CtaTwo from "../../cta/CtaTwo";
import FooterTwo from "../../../layout/Footer/FooterTwo";

const SecondBlog = () => {
  return (
    <Layout>
      <PageMeta title="Blog Details - Software &amp; IT Solutions HTML Template" />
      <Navbar navDark />
      <PageHeader title="Talent Acquisition Trends 2025" />
      <BlogDetailsSecond />
      {/* <NewsLetter /> */}
      <LatestBlog />
      <CtaTwo />
      <FooterTwo />
    </Layout>
  );
};

export default SecondBlog;

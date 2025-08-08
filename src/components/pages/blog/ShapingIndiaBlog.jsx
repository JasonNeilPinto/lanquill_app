import React from "react";
import Layout from "../../../layout/Layout";
import PageMeta from "../../common/PageMeta";
import Navbar from "../../../layout/Header/Navbar";
import PageHeader from "../../common/PageHeader";
import LatestBlog from "../../blogs/LatestBlog";
import CtaTwo from "../../cta/CtaTwo";
import ShapingIndia from "../../blogs/ShapingIndia";
import FooterTwo from "../../../layout/Footer/FooterTwo";

const ShapingIndiaBlog = () => {
  return (
    <Layout>
      <PageMeta title="Blog Details - Software &amp; IT Solutions HTML Template" />
      <Navbar navDark />
      <PageHeader title="Shaping India's Future: 12 Big Ideas for 2025 " />
      <ShapingIndia />
      {/* <NewsLetter /> */}
      <LatestBlog />
      <CtaTwo />
      <FooterTwo />
    </Layout>
  );
};

export default ShapingIndiaBlog;

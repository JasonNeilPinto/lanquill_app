import React from "react";
import AboutPageHero from "../../about/AboutPageHero";
import OurStory from "../../about/OurStory";
import FeatureImgThree from "../../features/FeatureImgThree";
import Team from "../../team/Team";
import TestimonialTwo from "../../testimonials/TestimonialTwo";
import PageMeta from "../../common/PageMeta";
import CtaTwo from "../../cta/CtaTwo";
import FooterOne from "../../../layout/Footer/FooterOne";
import Navbar from "../../../layout/Header/Navbar";
import Layout from "../../../layout/Layout";

const About = () => {
  return (
    <Layout>
      <PageMeta title="About us - Software &amp; IT Solutions HTML Template" />
      <Navbar classOption="navbar-light" />
      <AboutPageHero />
      <OurStory />
      <FeatureImgThree />
      <Team />
      <TestimonialTwo />
      <CtaTwo />
      <FooterOne footerLight />
    </Layout>
  );
};

export default About;

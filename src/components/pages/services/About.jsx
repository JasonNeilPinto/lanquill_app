import React from "react";
import AboutPageHero from "../../about/AboutPageHero";
import OurStory from "../../about/OurStory";
import TestimonialTwo from "../../testimonials/TestimonialTwo";
import PageMeta from "../../common/PageMeta";
import CtaTwo from "../../cta/CtaTwo";
import Navbar from "../../../layout/Header/Navbar";
import Layout from "../../../layout/Layout";
import Values from "../../about/AboutValues";
import FooterTwo from "../../../layout/Footer/FooterTwo";
import Journey from "../../others/Journey";

const About = () => {
  return (
    <Layout>
      <PageMeta title="About us - Software &amp; IT Solutions HTML Template" />
      <Navbar classOption="navbar-light" />
      <AboutPageHero />
      <Journey />
      <OurStory />
      <Values />
      <TestimonialTwo />
      <CtaTwo />
      <FooterTwo />
    </Layout>
  );
};

export default About;

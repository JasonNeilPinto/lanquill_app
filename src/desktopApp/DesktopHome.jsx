import React from "react";
import HeroTitle from "../components/common/HeroTitle.jsx";
import { Swiper, SwiperSlide } from "swiper/react";
import { Autoplay, Pagination } from "swiper";
import "swiper/css";
import "swiper/css/pagination";
import { Link } from "react-router-dom";

const slides = [
  {
    title: "Data & Analytics Solution",
    subtitle: "Build, Transform, Scale - Your Data, Our Expertise",
    desc: "We build, transform, and scale data systems to turn your information into impact.",
    macImg: "/img/screen/mac-screen/DataAnalytics.png",
    phoneImg: "/img/screen/phone-screen/dataAnalytics.png",
    link: "/services/data-analytics",
  },
  {
    title: "Cyber Security Solution",
    subtitle: "Future-Ready Security for Cloud and AI Systems",
    desc: "Delivering smart security solutions that scale with your business. ",
    macImg: "/img/screen/mac-screen/Cybersecurity.png",
    phoneImg: "/img/screen/phone-screen/cybersecurity.png",
    link: "/services/cyber-security",
  },
  {
    title: "Cloud Services Solution",
    subtitle: "From Legacy Limits to Cloud Possibilities",
    desc: "Reimagine your apps for a smarter, leaner cloud future.",
    macImg: "/img/screen/mac-screen/CloudServices.png",
    phoneImg: "/img/screen/phone-screen/cloudServices.png",
    link: "/services/cloud-services",
  },
  {
    title: "Applied AI Solution",
    subtitle: "Gen AI Solutions for Real Business Needs",
    desc: "We turn business needs into Gen AI-powered solutions - efficient, smart, and business-ready.",
    macImg: "/img/screen/mac-screen/AppliedAI.png",
    phoneImg: "/img/screen/phone-screen/appliedAI.png",
    link: "/services/applied-ai",
  },
  {
    title: "Security testing",
    subtitle: "Ensuring your AI solutions are safe, resilient, and compliant",
    desc: "We turn business needs into Gen AI-powered solutions - efficient, smart, and business-ready.",
    macImg: "/img/screen/mac-screen/SecurityTesting.png",
    phoneImg: "/img/screen/phone-screen/securityTesting.png",
    link: "/services/security-testing",
  },
  {
    title: "Managed Services",
    subtitle: "Reliable oversight for smarter AI solutions",
    desc: "We turn business needs into Gen AI-powered solutions - efficient, smart, and business-ready.",
    macImg: "/img/screen/mac-screen/ManagedServices.png",
    phoneImg: "/img/screen/phone-screen/managedServices.png",
    link: "/services/managed-services",
  },
  {
    title: "PramitiHR.AI",
    subtitle: "Hire Smarter with Generative AI Interviews",
    desc: "Revolutionizing Recruitment with AI-Driven Interviews",
    macImg: "/img/screen/mac-screen/PramitiHR.png",
    phoneImg: "/img/screen/phone-screen/pramitiHR.png",
    link: "#pramitihr",
  },
];

const DesktopHome = () => {
  return (
    <>
      <Swiper
        modules={[Autoplay, Pagination]}
        autoplay={{ delay: 2000 }}
        pagination={{ clickable: true }}
        loop={true}
        className="w-full"
      >
        {slides.map((slide, index) => (
          <SwiperSlide key={index}>
            <section
              className="hero-section ptb-120"
              style={{
                background:
                  "url('/img/shape/dot-dot-wave-shape.svg') no-repeat bottom center",
              }}
            >
              <div className="container">
                <div className="row align-items-center justify-content-lg-between">
                  <div className="col-xl-5 col-lg-5">
                    <div
                      className="hero-content-wrap text-center text-xl-start text-lg-start"
                      data-aos="fade-right"
                    >
                      <HeroTitle
                        title={slide.title}
                        desc={slide.desc}
                        subtitle={slide.subtitle}
                      />

                      <div className="pt-4 text-center text-xl-start text-lg-start">
                        <Link to={slide.link} className="btn btn-primary">
                          View More
                        </Link>
                      </div>
                    </div>
                  </div>

                  <div className="col-xl-6 col-lg-6 mt-4 mt-xl-0">
                    <div
                      className="hero-img-wrap position-relative"
                      data-aos="fade-left"
                    >
                      <ul className="position-absolute animate-element parallax-element shape-service hide-medium">
                        <li className="layer" data-depth="0.03">
                          <img
                            src="/img/color-shape/image-1.svg"
                            alt="shape"
                            className="img-fluid position-absolute color-shape-1"
                          />
                        </li>
                        <li className="layer" data-depth="0.03">
                          <img
                            src="/img/color-shape/feature-3.svg"
                            alt="shape"
                            className="img-fluid position-absolute color-shape-3"
                          />
                        </li>
                      </ul>

                      <div className="hero-img-wrap position-relative">
                        <div className="hero-screen-wrap">
                          <div className="phone-screen">
                            <img
                              src={slide.phoneImg}
                              alt="hero phone"
                              className="position-relative img-fluid"
                            />
                          </div>
                          <div className="mac-screen">
                            <img
                              src={slide.macImg}
                              alt="hero mac"
                              className="position-relative img-fluid rounded-custom"
                            />
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </section>
          </SwiperSlide>
        ))}
      </Swiper>
    </>
  );
};

export default DesktopHome;

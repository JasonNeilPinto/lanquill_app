/* eslint-disable react/prop-types */
import React from "react";
import SectionTitle from "../../common/SectionTitle";

const Feature = ({ cardDark }) => {
  return (
    <>
      <section
        className={`feature-section ptb-120 ${
          cardDark ? "bg-dark" : "bg-light"
        }`}
      >
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-lg-10 col-md-10">
              {cardDark ? (
                <SectionTitle
                  subtitle="Services"
                  title="Best Services Grow Your Business Value"
                  description="Globally actualize cost effective with resource maximizing
                  leadership skills."
                  centerAlign
                  dark
                />
              ) : (
                <SectionTitle
                  subtitle="Services"
                  title="Explore our Services"
                  description="From intelligent automation to predictive analytics and synthetic content creation, we empower enterprises to stay ahead in an AI-first world."
                  centerAlign
                />
              )}
            </div>
          </div>
          <div className="row">
            <div className="col-12">
              <div className="feature-grid">
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="far fa-database icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Generate</h3>
                    <p className="mb-0">
                      Leverage generative AI for automated content creation,
                      synthetic media, conversational AI, and code
                      generation—accelerating creativity across marketing,
                      support, and development.
                    </p>
                  </div>
                </div>
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="fal fa-lightbulb icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Action</h3>
                    <p className="mb-0">
                      Enable real-time efficiency with autonomous agents and
                      intelligent automation that reduce manual work and
                      optimize workflows.
                    </p>
                  </div>
                </div>
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="fal fa-chart-line icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Linguistics</h3>
                    <p className="mb-0">
                      Enhance communication with{" "}
                      <b>Natural Language Processing</b>
                      (NLP), language understanding, and cross-lingual AI for
                      chatbots, voice assistants, and multilingual support.
                    </p>
                  </div>
                </div>
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="fal fa-file-chart-line icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Optics</h3>
                    <p className="mb-0">
                      Extract insights from visuals through image recognition,
                      video analytics, AI-powered medical imaging, and quality
                      inspection automation.
                    </p>
                  </div>
                </div>
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="fal fa-exchange-alt icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Pattern</h3>
                    <p className="mb-0">
                      Drive data-backed decisions with predictive analytics,
                      recommendation systems, anomaly detection, and risk
                      scoring for personalization and fraud prevention.
                    </p>
                  </div>
                </div>
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="fal fa-compass icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Sonics</h3>
                    <p className="mb-0">
                      Tap into voice data using audio classification, speech
                      analytics, and speaker diarization—ideal for contact
                      centers and voice-driven apps.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default Feature;

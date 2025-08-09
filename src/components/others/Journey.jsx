import React from "react";
import SectionTitle from "../common/SectionTitle";

const Journey = () => {
  return (
    <>
      <section className="work-process ptb-120">
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-lg-6 col-md-10">
              <SectionTitle
                subtitle="Our Journey"
                title="A Timeline of Growth and Innovation"
                description="From startup to global IT consulting leader"
                centerAlign
              />
            </div>
          </div>
          <div className="row align-items-center justify-content-between">
            <div className="col-lg-12 col-md-12 order-0 order-lg-1">
              <div className="timeline-wrapper position-relative mt-5">
                {[
                  {
                    step: "2015",
                    icon: "far fa-lightbulb",
                    title: "THE START",
                    desc: "Started with digitization and retail analytics products. Gained expertise in SMAC technologies and started IT Consulting Services in Cyber Security.",
                  },
                  {
                    step: "2018",
                    icon: "far fa-rocket",
                    title: "PRODUCT LAUNCH",
                    desc: "Launched products in digitization/infotainment, started building AI/ML product, and began managed services in IT Infrastructure Management and Cyber Security.",
                  },
                  {
                    step: "2020",
                    icon: "far fa-globe-americas",
                    title: "EXPANSION",
                    desc: "Started operations in Australia and diversified into custom application development, DevOps and cloud automation services. Enabled remote work capabilities.",
                  },
                  {
                    step: "2023",
                    icon: "far fa-chart-line",
                    title: "GROWTH",
                    desc: "Building CoE, Tools and Accelerators in Custom Application Development, Application Security, VAPT, DevSecOps and Cloud Automation. Positioned to reach 100+ consultants globally.",
                  },
                ].map((item, idx) => (
                  <div
                    key={idx}
                    className={`timeline-step ${
                      idx % 2 === 0 ? "left" : "right"
                    }`}
                  >
                    <div className="content bg-white rounded-custom custom-shadow p-3">
                      <div className="d-flex align-items-center mb-2">
                        <h5 className="text-primary me-4">
                          {item.step}
                        </h5>
                        <span className="badge bg-secondary mb-2">
                          {item.title}
                        </span>
                      </div>
                      <p>{item.desc}</p>
                    </div>
                    <div
                      className={`${
                        idx === 3 ? "process-icon-1" : "process-icon-2"
                      } border border-2 rounded-custom bg-white me-4 mt-2 icon`}
                    >
                      <i className={`${item.icon} fa-2x text-primary`}></i>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default Journey;

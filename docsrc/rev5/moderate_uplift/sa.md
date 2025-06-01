# SA - System and Services Acquisition

* Controls: 8

## Controls

### SA-4 (1): Functional Properties of Controls

Require the developer of the system, system component, or system service to provide a description of the functional properties of the controls to be implemented.

Functional properties of security and privacy controls describe the functionality (i.e., security or privacy capability, functions, or mechanisms) visible at the interfaces of the controls and specifically exclude functionality and data structures internal to the operation of the controls.

the developer of the system, system component, or system service is required to provide a description of the functional properties of the controls to be implemented.

System and services acquisition policy

system and services acquisition procedures

procedures addressing the integration of security and privacy requirements, descriptions, and criteria into the acquisition process

solicitation documents

acquisition documentation

acquisition contracts for the system, system component, or system services

system security plan

privacy plan

other relevant documents or records

Organizational personnel with acquisition/contracting responsibilities

organizational personnel with information security and privacy responsibilities

system developers

Organizational processes for determining system security functional requirements

organizational processes for developing acquisition contracts

mechanisms supporting and/or implementing acquisitions and the inclusion of security and privacy requirements in contracts

### SA-4 (2): Design and Implementation Information for Controls

Require the developer of the system, system component, or system service to provide design and implementation information for the controls that includes: security-relevant external system interfaces, high-level design, low-level design, source code or hardware schematics, and/or  design and implementation information is defined (if selected); at level of detail is defined;.

Organizations may require different levels of detail in the documentation for the design and implementation of controls in organizational systems, system components, or system services based on mission and business requirements, requirements for resiliency and trustworthiness, and requirements for analysis and testing. Systems can be partitioned into multiple subsystems. Each subsystem within the system can contain one or more modules. The high-level design for the system is expressed in terms of subsystems and the interfaces between subsystems providing security-relevant functionality. The low-level design for the system is expressed in terms of modules and the interfaces between modules providing security-relevant functionality. Design and implementation documentation can include manufacturer, version, serial number, verification hash signature, software libraries used, date of purchase or download, and the vendor or download source. Source code and hardware schematics are referred to as the implementation representation of the system.

the developer of the system, system component, or system service is required to provide design and implementation information for the controls that includes using security-relevant external system interfaces, high-level design, low-level design, source code or hardware schematics, and/or  design and implementation information is defined (if selected); at level of detail is defined;.

System and services acquisition policy

system and services acquisition procedures

procedures addressing the integration of security requirements, descriptions, and criteria into the acquisition process

solicitation documents

acquisition documentation

acquisition contracts for the system, system components, or system services

design and implementation information for controls employed in the system, system component, or system service

system security plan

other relevant documents or records

Organizational personnel with acquisition/contracting responsibilities

organizational personnel with the responsibility to determine system security requirements

system developers or service provider

organizational personnel with information security responsibilities

Organizational processes for determining the level of detail for system design and controls

organizational processes for developing acquisition contracts

mechanisms supporting and/or implementing the development of system design details

### SA-4 (9): Functions, Ports, Protocols, and Services in Use

Require the developer of the system, system component, or system service to identify the functions, ports, protocols, and services intended for organizational use.

The identification of functions, ports, protocols, and services early in the system development life cycle (e.g., during the initial requirements definition and design stages) allows organizations to influence the design of the system, system component, or system service. This early involvement in the system development life cycle helps organizations avoid or minimize the use of functions, ports, protocols, or services that pose unnecessarily high risks and understand the trade-offs involved in blocking specific ports, protocols, or services or requiring system service providers to do so. Early identification of functions, ports, protocols, and services avoids costly retrofitting of controls after the system, component, or system service has been implemented. [SA-9](#sa-9) describes the requirements for external system services. Organizations identify which functions, ports, protocols, and services are provided from external sources.

the developer of the system, system component, or system service is required to identify the functions intended for organizational use;

the developer of the system, system component, or system service is required to identify the ports intended for organizational use;

the developer of the system, system component, or system service is required to identify the protocols intended for organizational use;

the developer of the system, system component, or system service is required to identify the services intended for organizational use.

System and services acquisition policy

procedures addressing the integration of security requirements, descriptions, and criteria into the acquisition process

system design documentation

system documentation, including functions, ports, protocols, and services intended for organizational use

acquisition contracts for systems or services

acquisition documentation

solicitation documentation

service level agreements

organizational security requirements, descriptions, and criteria for developers of systems, system components, and system services

system security plan

other relevant documents or records

Organizational personnel with acquisition/contracting responsibilities

organizational personnel with the responsibility for determining system security requirements

system/network administrators

organizational personnel operating, using, and/or maintaining the system

system developers

organizational personnel with information security responsibilities

### SA-9 (2): Identification of Functions, Ports, Protocols, and Services

Require providers of the following external system services to identify the functions, ports, protocols, and other services required for the use of such services: external system services that require the identification of functions, ports, protocols, and other services are defined;.

Information from external service providers regarding the specific functions, ports, protocols, and services used in the provision of such services can be useful when the need arises to understand the trade-offs involved in restricting certain functions and services or blocking certain ports and protocols.

providers of external system services that require the identification of functions, ports, protocols, and other services are defined; are required to identify the functions, ports, protocols, and other services required for the use of such services.

System and services acquisition policy

supply chain risk management policy and procedures

procedures addressing external system services

acquisition contracts for the system, system component, or system service

acquisition documentation

solicitation documentation

service level agreements

organizational security requirements and security specifications for external service providers

list of required functions, ports, protocols, and other services

system security plan

other relevant documents or records

Organizational personnel with system and service acquisition responsibilities

organizational personnel with information security responsibilities

system/network administrators

external providers of system services

### SA-10: Developer Configuration Management

Require the developer of the system, system component, or system service to:

Perform configuration management during system, component, or service design, development, implementation, operation, and/or disposal;

Document, manage, and control the integrity of changes to configuration items under configuration management are defined;;

Implement only organization-approved changes to the system, component, or service;

Document approved changes to the system, component, or service and the potential security and privacy impacts of such changes; and

Track security flaws and flaw resolution within the system, component, or service and report findings to personnel to whom security flaws and flaw resolutions within the system, component, or service are reported is/are defined;.

Organizations consider the quality and completeness of configuration management activities conducted by developers as direct evidence of applying effective security controls. Controls include protecting the master copies of material used to generate security-relevant portions of the system hardware, software, and firmware from unauthorized modification or destruction. Maintaining the integrity of changes to the system, system component, or system service requires strict configuration control throughout the system development life cycle to track authorized changes and prevent unauthorized changes.

The configuration items that are placed under configuration management include the formal model; the functional, high-level, and low-level design specifications; other design data; implementation documentation; source code and hardware schematics; the current running version of the object code; tools for comparing new versions of security-relevant hardware descriptions and source code with previous versions; and test fixtures and documentation. Depending on the mission and business needs of organizations and the nature of the contractual relationships in place, developers may provide configuration management support during the operations and maintenance stage of the system development life cycle.

the developer of the system, system component, or system service is required to perform configuration management during system, component, or service design, development, implementation, operation, and/or disposal;

the developer of the system, system component, or system service is required to document the integrity of changes to configuration items under configuration management are defined;;

the developer of the system, system component, or system service is required to manage the integrity of changes to configuration items under configuration management are defined;;

the developer of the system, system component, or system service is required to control the integrity of changes to configuration items under configuration management are defined;;

the developer of the system, system component, or system service is required to implement only organization-approved changes to the system, component, or service;

the developer of the system, system component, or system service is required to document approved changes to the system, component, or service;

the developer of the system, system component, or system service is required to document the potential security impacts of approved changes;

the developer of the system, system component, or system service is required to document the potential privacy impacts of approved changes;

the developer of the system, system component, or system service is required to track security flaws within the system, component, or service;

the developer of the system, system component, or system service is required to track security flaw resolutions within the system, component, or service;

the developer of the system, system component, or system service is required to report findings to personnel to whom security flaws and flaw resolutions within the system, component, or service are reported is/are defined;.

System and services acquisition policy

procedures addressing system developer configuration management

solicitation documentation

acquisition documentation

service level agreements

acquisition contracts for the system, system component, or system service

system developer configuration management plan

security flaw and flaw resolution tracking records

system change authorization records

change control records

configuration management records

system security plan

other relevant documents or records

Organizational personnel with system and service acquisition responsibilities

organizational personnel with information security responsibilities

organizational personnel with configuration management responsibilities

system developers

Organizational processes for monitoring developer configuration management

mechanisms supporting and/or implementing the monitoring of developer configuration management

### SA-11: Developer Testing and Evaluation

Require the developer of the system, system component, or system service, at all post-design stages of the system development life cycle, to:

Develop and implement a plan for ongoing security and privacy control assessments;

Perform unit, integration, system, and/or regression testing/evaluation frequency at which to conduct unit, integration, system, and/or regression testing/evaluation is defined; at depth and coverage of unit, integration, system, and/or regression testing/evaluation is defined;;

Produce evidence of the execution of the assessment plan and the results of the testing and evaluation;

Implement a verifiable flaw remediation process; and

Correct flaws identified during testing and evaluation.

Developmental testing and evaluation confirms that the required controls are implemented correctly, operating as intended, enforcing the desired security and privacy policies, and meeting established security and privacy requirements. Security properties of systems and the privacy of individuals may be affected by the interconnection of system components or changes to those components. The interconnections or changes—including upgrading or replacing applications, operating systems, and firmware—may adversely affect previously implemented controls. Ongoing assessment during development allows for additional types of testing and evaluation that developers can conduct to reduce or eliminate potential flaws. Testing custom software applications may require approaches such as manual code review, security architecture review, and penetration testing, as well as and static analysis, dynamic analysis, binary analysis, or a hybrid of the three analysis approaches.

Developers can use the analysis approaches, along with security instrumentation and fuzzing, in a variety of tools and in source code reviews. The security and privacy assessment plans include the specific activities that developers plan to carry out, including the types of analyses, testing, evaluation, and reviews of software and firmware components; the degree of rigor to be applied; the frequency of the ongoing testing and evaluation; and the types of artifacts produced during those processes. The depth of testing and evaluation refers to the rigor and level of detail associated with the assessment process. The coverage of testing and evaluation refers to the scope (i.e., number and type) of the artifacts included in the assessment process. Contracts specify the acceptance criteria for security and privacy assessment plans, flaw remediation processes, and the evidence that the plans and processes have been diligently applied. Methods for reviewing and protecting assessment plans, evidence, and documentation are commensurate with the security category or classification level of the system. Contracts may specify protection requirements for documentation.

the developer of the system, system component, or system service is required at all post-design stages of the system development life cycle to develop a plan for ongoing security assessments;

the developer of the system, system component, or system service is required at all post-design stages of the system development life cycle to implement a plan for ongoing security assessments;

the developer of the system, system component, or system service is required at all post-design stages of the system development life cycle to develop a plan for privacy assessments;

the developer of the system, system component, or system service is required at all post-design stages of the system development life cycle to implement a plan for ongoing privacy assessments;

the developer of the system, system component, or system service is required at all post-design stages of the system development life cycle to perform unit, integration, system, and/or regression testing/evaluation frequency at which to conduct unit, integration, system, and/or regression testing/evaluation is defined; at depth and coverage of unit, integration, system, and/or regression testing/evaluation is defined;;

the developer of the system, system component, or system service is required at all post-design stages of the system development life cycle to produce evidence of the execution of the assessment plan;

the developer of the system, system component, or system service is required at all post-design stages of the system development life cycle to produce the results of the testing and evaluation;

the developer of the system, system component, or system service is required at all post-design stages of the system development life cycle to implement a verifiable flaw remediation process;

the developer of the system, system component, or system service is required at all post-design stages of the system development life cycle to correct flaws identified during testing and evaluation.

System and services acquisition policy

system and services acquisition procedures

procedures addressing system developer security and privacy testing

procedures addressing flaw remediation

solicitation documentation

acquisition documentation

service level agreements

acquisition contracts for the system, system component, or system service

security and privacy architecture

system design documentation

system developer security and privacy assessment plans

results of developer security and privacy assessments for the system, system component, or system service

security and privacy flaw and remediation tracking records

system security plan

privacy plan

privacy impact assessment

privacy risk assessment documentation

other relevant documents or records

Organizational personnel with system and service acquisition responsibilities

organizational personnel with information security and privacy responsibilities

organizational personnel with developer security and privacy testing responsibilities

system developers

Organizational processes for monitoring developer security testing and evaluation

mechanisms supporting and/or implementing the monitoring of developer security and privacy testing and evaluation

### SA-15: Development Process, Standards, and Tools

Require the developer of the system, system component, or system service to follow a documented development process that:

Explicitly addresses security and privacy requirements;

Identifies the standards and tools used in the development process;

Documents the specific tool options and tool configurations used in the development process; and

Documents, manages, and ensures the integrity of changes to the process and/or tools used in development; and

Review the development process, standards, tools, tool options, and tool configurations frequency at which to review the development process, standards, tools, tool options, and tool configurations is defined; to determine if the process, standards, tools, tool options and tool configurations selected and employed can satisfy the following security and privacy requirements: organization-defined security and privacy requirements.

Development tools include programming languages and computer-aided design systems. Reviews of development processes include the use of maturity models to determine the potential effectiveness of such processes. Maintaining the integrity of changes to tools and processes facilitates effective supply chain risk assessment and mitigation. Such integrity requires configuration control throughout the system development life cycle to track authorized changes and prevent unauthorized changes.

the developer of the system, system component, or system service is required to follow a documented development process that explicitly addresses security requirements;

the developer of the system, system component, or system service is required to follow a documented development process that explicitly addresses privacy requirements;

the developer of the system, system component, or system service is required to follow a documented development process that identifies the standards used in the development process;

the developer of the system, system component, or system service is required to follow a documented development process that identifies the tools used in the development process;

the developer of the system, system component, or system service is required to follow a documented development process that documents the specific tool used in the development process;

the developer of the system, system component, or system service is required to follow a documented development process that documents the specific tool configurations used in the development process;

the developer of the system, system component, or system service is required to follow a documented development process that documents, manages, and ensures the integrity of changes to the process and/or tools used in development;

the developer of the system, system component, or system service is required to follow a documented development process in which the development process, standards, tools, tool options, and tool configurations are reviewed frequency at which to review the development process, standards, tools, tool options, and tool configurations is defined; to determine that the process, standards, tools, tool options, and tool configurations selected and employed satisfy security requirements to be satisfied by the process, standards, tools, tool options, and tool configurations are defined;;

the developer of the system, system component, or system service is required to follow a documented development process in which the development process, standards, tools, tool options, and tool configurations are reviewed frequency at which to review the development process, standards, tools, tool options, and tool configurations is defined; to determine that the process, standards, tools, tool options, and tool configurations selected and employed satisfy privacy requirements to be satisfied by the process, standards, tools, tool options, and tool configurations are defined;.

System and services acquisition policy

system and services acquisition procedures

procedures addressing development process, standards, and tools

procedures addressing the integration of security and privacy requirements during the development process

solicitation documentation

acquisition documentation

critical component inventory documentation

service level agreements

acquisition contracts for the system, system component, or system service

system developer documentation listing tool options/configuration guides

configuration management policy

configuration management records

documentation of development process reviews using maturity models

change control records

configuration control records

documented reviews of the development process, standards, tools, and tool options/configurations

system security plan

privacy plan

privacy impact assessment

privacy risk assessment documentation

other relevant documents or records

Organizational personnel with system and service acquisition responsibilities

organizational personnel with information security and privacy responsibilities

system developer

### SA-15 (3): Criticality Analysis

Require the developer of the system, system component, or system service to perform a criticality analysis:

At the following decision points in the system development life cycle: decision points in the system development life cycle are defined; ; and

At the following level of rigor: organization-defined breadth and depth of criticality analysis.

Criticality analysis performed by the developer provides input to the criticality analysis performed by organizations. Developer input is essential to organizational criticality analysis because organizations may not have access to detailed design documentation for system components that are developed as commercial off-the-shelf products. Such design documentation includes functional specifications, high-level designs, low-level designs, source code, and hardware schematics. Criticality analysis is important for organizational systems that are designated as high value assets. High value assets can be moderate- or high-impact systems due to heightened adversarial interest or potential adverse effects on the federal enterprise. Developer input is especially important when organizations conduct supply chain criticality analyses.

the developer of the system, system component, or system service is required to perform a criticality analysis at decision points in the system development life cycle are defined; in the system development life cycle;

the developer of the system, system component, or system service is required to perform a criticality analysis at the following rigor level: the breadth of criticality analysis is defined;;

the developer of the system, system component, or system service is required to perform a criticality analysis at the following rigor level: the depth of criticality analysis is defined; .

Supply chain risk management plan

system and services acquisition policy

procedures addressing development process, standards, and tools

procedures addressing criticality analysis requirements for the system, system component, or system service

solicitation documentation

acquisition documentation

service level agreements

acquisition contracts for the system, system component, or system service

criticality analysis documentation

business impact analysis documentation

software development life cycle documentation

system security plan

other relevant documents or records

Organizational personnel with system and service acquisition responsibilities

organizational personnel with information security responsibilities

organizational personnel responsible for performing criticality analysis

system developer

organizational personnel with supply chain risk management responsibilities

Organizational processes for performing criticality analysis

mechanisms supporting and/or implementing criticality analysis


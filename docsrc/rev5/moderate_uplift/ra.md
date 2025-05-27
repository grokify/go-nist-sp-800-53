# RA - Risk Assessment

* Controls: 2

## Controls

### RA-5 (5): Privileged Access

Implement privileged access authorization to {{ insert: param, ra-05.05_odp.01 }} for {{ insert: param, ra-05.05_odp.02 }}.

In certain situations, the nature of the vulnerability scanning may be more intrusive, or the system component that is the subject of the scanning may contain classified or controlled unclassified information, such as personally identifiable information. Privileged access authorization to selected system components facilitates more thorough vulnerability scanning and protects the sensitive nature of such scanning.

privileged access authorization is implemented to {{ insert: param, ra-05.05_odp.01 }} for {{ insert: param, ra-05.05_odp.02 }}.

Risk assessment policy

procedures addressing vulnerability scanning

system design documentation

system configuration settings and associated documentation

list of system components for vulnerability scanning

personnel access authorization list

authorization credentials

access authorization records

system security plan

other relevant documents or records

Organizational personnel with vulnerability scanning responsibilities

system/network administrators

organizational personnel responsible for access control to the system

organizational personnel responsible for configuration management of the system

system developers

organizational personnel with security responsibilities

Organizational processes for vulnerability scanning

organizational processes for access control

mechanisms supporting and/or implementing access control

mechanisms/tools supporting and/or implementing vulnerability scanning

### RA-9: Criticality Analysis

Identify critical system components and functions by performing a criticality analysis for {{ insert: param, ra-09_odp.01 }} at {{ insert: param, ra-09_odp.02 }}.

Not all system components, functions, or services necessarily require significant protections. For example, criticality analysis is a key tenet of supply chain risk management and informs the prioritization of protection activities. The identification of critical system components and functions considers applicable laws, executive orders, regulations, directives, policies, standards, system functionality requirements, system and component interfaces, and system and component dependencies. Systems engineers conduct a functional decomposition of a system to identify mission-critical functions and components. The functional decomposition includes the identification of organizational missions supported by the system, decomposition into the specific functions to perform those missions, and traceability to the hardware, software, and firmware components that implement those functions, including when the functions are shared by many components within and external to the system.

The operational environment of a system or a system component may impact the criticality, including the connections to and dependencies on cyber-physical systems, devices, system-of-systems, and outsourced IT services. System components that allow unmediated access to critical system components or functions are considered critical due to the inherent vulnerabilities that such components create. Component and function criticality are assessed in terms of the impact of a component or function failure on the organizational missions that are supported by the system that contains the components and functions.

Criticality analysis is performed when an architecture or design is being developed, modified, or upgraded. If such analysis is performed early in the system development life cycle, organizations may be able to modify the system design to reduce the critical nature of these components and functions, such as by adding redundancy or alternate paths into the system design. Criticality analysis can also influence the protection measures required by development contractors. In addition to criticality analysis for systems, system components, and system services, criticality analysis of information is an important consideration. Such analysis is conducted as part of security categorization in [RA-2](#ra-2).

critical system components and functions are identified by performing a criticality analysis for {{ insert: param, ra-09_odp.01 }} at {{ insert: param, ra-09_odp.02 }}.

Risk assessment policy

assessment reports

criticality analysis/finalized criticality for each component/subcomponent

audit records/event logs

analysis reports

system security plan

other relevant documents or records

Organizational personnel with assessment and auditing responsibilities

organizational personnel with criticality analysis responsibilities

system/network administrators

organizational personnel with security responsibilities

Organizational processes for assessments and audits

mechanisms/tools supporting and/or implementing assessments and auditing


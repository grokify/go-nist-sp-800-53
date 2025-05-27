# ma - Maintenance

* Controls: 4

## Controls

### {ma-01 ma %!s(int=1) %!s(int=0)}: Policy and Procedures

Develop, document, and disseminate to {{ insert: param, ma-1_prm_1 }}:

 {{ insert: param, ma-01_odp.03 }} maintenance policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the maintenance policy and the associated maintenance controls;

Designate an {{ insert: param, ma-01_odp.04 }} to manage the development, documentation, and dissemination of the maintenance policy and procedures; and

Review and update the current maintenance:

Policy {{ insert: param, ma-01_odp.05 }} and following {{ insert: param, ma-01_odp.06 }} ; and

Procedures {{ insert: param, ma-01_odp.07 }} and following {{ insert: param, ma-01_odp.08 }}.

Maintenance policy and procedures address the controls in the MA family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of maintenance policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies that reflect the complex nature of organizations. Procedures can be established for security and privacy programs, for mission or business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to maintenance policy and procedures assessment or audit findings, security incidents or breaches, or changes in applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

a maintenance policy is developed and documented;

the maintenance policy is disseminated to {{ insert: param, ma-01_odp.01 }};

maintenance procedures to facilitate the implementation of the maintenance policy and associated maintenance controls are developed and documented;

the maintenance procedures are disseminated to {{ insert: param, ma-01_odp.02 }};

the {{ insert: param, ma-01_odp.03 }} maintenance policy addresses purpose;

the {{ insert: param, ma-01_odp.03 }} maintenance policy addresses scope;

the {{ insert: param, ma-01_odp.03 }} maintenance policy addresses roles;

the {{ insert: param, ma-01_odp.03 }} maintenance policy addresses responsibilities;

the {{ insert: param, ma-01_odp.03 }} maintenance policy addresses management commitment;

the {{ insert: param, ma-01_odp.03 }} maintenance policy addresses coordination among organizational entities;

the {{ insert: param, ma-01_odp.03 }} maintenance policy addresses compliance;

the {{ insert: param, ma-01_odp.03 }} maintenance policy is consistent with applicable laws, Executive Orders, directives, regulations, policies, standards, and guidelines;

the {{ insert: param, ma-01_odp.04 }} is designated to manage the development, documentation, and dissemination of the maintenance policy and procedures;

the current maintenance policy is reviewed and updated {{ insert: param, ma-01_odp.05 }};

the current maintenance policy is reviewed and updated following {{ insert: param, ma-01_odp.06 }};

the current maintenance procedures are reviewed and updated {{ insert: param, ma-01_odp.07 }};

the current maintenance procedures are reviewed and updated following {{ insert: param, ma-01_odp.08 }}.

Maintenance policy and procedures

system security plan

privacy plan

organizational risk management strategy

other relevant documents or records

Organizational personnel with maintenance responsibilities

organizational personnel with information security and privacy responsibilities

### {ma-02 ma %!s(int=2) %!s(int=0)}: Controlled Maintenance

Schedule, document, and review records of maintenance, repair, and replacement on system components in accordance with manufacturer or vendor specifications and/or organizational requirements;

Approve and monitor all maintenance activities, whether performed on site or remotely and whether the system or system components are serviced on site or removed to another location;

Require that {{ insert: param, ma-02_odp.01 }} explicitly approve the removal of the system or system components from organizational facilities for off-site maintenance, repair, or replacement;

Sanitize equipment to remove the following information from associated media prior to removal from organizational facilities for off-site maintenance, repair, or replacement: {{ insert: param, ma-02_odp.02 }};

Check all potentially impacted controls to verify that the controls are still functioning properly following maintenance, repair, or replacement actions; and

Include the following information in organizational maintenance records: {{ insert: param, ma-02_odp.03 }}.

Controlling system maintenance addresses the information security aspects of the system maintenance program and applies to all types of maintenance to system components conducted by local or nonlocal entities. Maintenance includes peripherals such as scanners, copiers, and printers. Information necessary for creating effective maintenance records includes the date and time of maintenance, a description of the maintenance performed, names of the individuals or group performing the maintenance, name of the escort, and system components or equipment that are removed or replaced. Organizations consider supply chain-related risks associated with replacement components for systems.

maintenance, repair, and replacement of system components are scheduled in accordance with manufacturer or vendor specifications and/or organizational requirements;

maintenance, repair, and replacement of system components are documented in accordance with manufacturer or vendor specifications and/or organizational requirements;

records of maintenance, repair, and replacement of system components are reviewed in accordance with manufacturer or vendor specifications and/or organizational requirements;

all maintenance activities, whether performed on site or remotely and whether the system or system components are serviced on site or removed to another location, are approved;

all maintenance activities, whether performed on site or remotely and whether the system or system components are serviced on site or removed to another location, are monitored;

 {{ insert: param, ma-02_odp.01 }} is/are required to explicitly approve the removal of the system or system components from organizational facilities for off-site maintenance, repair, or replacement;

equipment is sanitized to remove {{ insert: param, ma-02_odp.02 }} from associated media prior to removal from organizational facilities for off-site maintenance, repair, or replacement;

all potentially impacted controls are checked to verify that the controls are still functioning properly following maintenance, repair, or replacement actions;

 {{ insert: param, ma-02_odp.03 }} is included in organizational maintenance records.

Maintenance policy

procedures addressing controlled system maintenance

maintenance records

manufacturer/vendor maintenance specifications

equipment sanitization records

media sanitization records

system security plan

other relevant documents or records

Organizational personnel with system maintenance responsibilities

organizational personnel with information security responsibilities

organizational personnel responsible for media sanitization

system/network administrators

Organizational processes for scheduling, performing, documenting, reviewing, approving, and monitoring maintenance and repairs for the system

organizational processes for sanitizing system components

mechanisms supporting and/or implementing controlled maintenance

mechanisms implementing the sanitization of system components

### {ma-04 ma %!s(int=4) %!s(int=0)}: Nonlocal Maintenance

Approve and monitor nonlocal maintenance and diagnostic activities;

Allow the use of nonlocal maintenance and diagnostic tools only as consistent with organizational policy and documented in the security plan for the system;

Employ strong authentication in the establishment of nonlocal maintenance and diagnostic sessions;

Maintain records for nonlocal maintenance and diagnostic activities; and

Terminate session and network connections when nonlocal maintenance is completed.

Nonlocal maintenance and diagnostic activities are conducted by individuals who communicate through either an external or internal network. Local maintenance and diagnostic activities are carried out by individuals who are physically present at the system location and not communicating across a network connection. Authentication techniques used to establish nonlocal maintenance and diagnostic sessions reflect the network access requirements in [IA-2](#ia-2) . Strong authentication requires authenticators that are resistant to replay attacks and employ multi-factor authentication. Strong authenticators include PKI where certificates are stored on a token protected by a password, passphrase, or biometric. Enforcing requirements in [MA-4](#ma-4) is accomplished, in part, by other controls. [SP 800-63B](#e59c5a7c-8b1f-49ca-8de0-6ee0882180ce) provides additional guidance on strong authentication and authenticators.

nonlocal maintenance and diagnostic activities are approved;

nonlocal maintenance and diagnostic activities are monitored;

the use of nonlocal maintenance and diagnostic tools are allowed only as consistent with organizational policy;

the use of nonlocal maintenance and diagnostic tools are documented in the security plan for the system;

strong authentication is employed in the establishment of nonlocal maintenance and diagnostic sessions;

records for nonlocal maintenance and diagnostic activities are maintained;

session connections are terminated when nonlocal maintenance is completed;

network connections are terminated when nonlocal maintenance is completed.

Maintenance policy

procedures addressing nonlocal system maintenance

remote access policy

remote access procedures

system design documentation

system configuration settings and associated documentation

maintenance records

records of remote access

diagnostic records

system security plan

other relevant documents or records

Organizational personnel with system maintenance responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing nonlocal maintenance

mechanisms implementing, supporting, and/or managing nonlocal maintenance

mechanisms for strong authentication of nonlocal maintenance diagnostic sessions

mechanisms for terminating nonlocal maintenance sessions and network connections

### {ma-05 ma %!s(int=5) %!s(int=0)}: Maintenance Personnel

Establish a process for maintenance personnel authorization and maintain a list of authorized maintenance organizations or personnel;

Verify that non-escorted personnel performing maintenance on the system possess the required access authorizations; and

Designate organizational personnel with required access authorizations and technical competence to supervise the maintenance activities of personnel who do not possess the required access authorizations.

Maintenance personnel refers to individuals who perform hardware or software maintenance on organizational systems, while [PE-2](#pe-2) addresses physical access for individuals whose maintenance duties place them within the physical protection perimeter of the systems. Technical competence of supervising individuals relates to the maintenance performed on the systems, while having required access authorizations refers to maintenance on and near the systems. Individuals not previously identified as authorized maintenance personnel—such as information technology manufacturers, vendors, systems integrators, and consultants—may require privileged access to organizational systems, such as when they are required to conduct maintenance activities with little or no notice. Based on organizational assessments of risk, organizations may issue temporary credentials to these individuals. Temporary credentials may be for one-time use or for very limited time periods.

a process for maintenance personnel authorization is established;

a list of authorized maintenance organizations or personnel is maintained;

non-escorted personnel performing maintenance on the system possess the required access authorizations;

organizational personnel with required access authorizations and technical competence is/are designated to supervise the maintenance activities of personnel who do not possess the required access authorizations.

Maintenance policy

procedures addressing maintenance personnel

service provider contracts

service-level agreements

list of authorized personnel

maintenance records

access control records

system security plan

other relevant documents or records

Organizational personnel with system maintenance responsibilities

organizational personnel with information security responsibilities

Organizational processes for authorizing and managing maintenance personnel

mechanisms supporting and/or implementing authorization of maintenance personnel


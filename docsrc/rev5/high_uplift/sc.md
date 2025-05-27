# SC - System and Communications Protection

* Controls: 5

## Controls

### SC-3: Security Function Isolation

Isolate security functions from nonsecurity functions.

Security functions are isolated from nonsecurity functions by means of an isolation boundary implemented within a system via partitions and domains. The isolation boundary controls access to and protects the integrity of the hardware, software, and firmware that perform system security functions. Systems implement code separation in many ways, such as through the provision of security kernels via processor rings or processor modes. For non-kernel code, security function isolation is often achieved through file system protections that protect the code on disk and address space protections that protect executing code. Systems can restrict access to security functions using access control mechanisms and by implementing least privilege capabilities. While the ideal is for all code within the defined security function isolation boundary to only contain security-relevant code, it is sometimes necessary to include nonsecurity functions as an exception. The isolation of security functions from nonsecurity functions can be achieved by applying the systems security engineering design principles in [SA-8](#sa-8) , including [SA-8(1)](#sa-8.1), [SA-8(3)](#sa-8.3), [SA-8(4)](#sa-8.4), [SA-8(10)](#sa-8.10), [SA-8(12)](#sa-8.12), [SA-8(13)](#sa-8.13), [SA-8(14)](#sa-8.14) , and [SA-8(18)](#sa-8.18).

security functions are isolated from non-security functions.

System and communications protection policy

procedures addressing security function isolation

list of security functions to be isolated from non-security functions

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

Separation of security functions from non-security functions within the system

### SC-7 (18): Fail Secure

Prevent systems from entering unsecure states in the event of an operational failure of a boundary protection device.

Fail secure is a condition achieved by employing mechanisms to ensure that in the event of operational failures of boundary protection devices at managed interfaces, systems do not enter into unsecure states where intended security properties no longer hold. Managed interfaces include routers, firewalls, and application gateways that reside on protected subnetworks (commonly referred to as demilitarized zones). Failures of boundary protection devices cannot lead to or cause information external to the devices to enter the devices nor can failures permit unauthorized information releases.

systems are prevented from entering unsecure states in the event of an operational failure of a boundary protection device.

System and communications protection policy

procedures addressing boundary protection

system design documentation

system architecture

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

organizational personnel with boundary protection responsibilities

Mechanisms supporting and/or implementing secure failure

### SC-7 (21): Isolation of System Components

Employ boundary protection mechanisms to isolate {{ insert: param, sc-07.21_odp.01 }} supporting {{ insert: param, sc-07.21_odp.02 }}.

Organizations can isolate system components that perform different mission or business functions. Such isolation limits unauthorized information flows among system components and provides the opportunity to deploy greater levels of protection for selected system components. Isolating system components with boundary protection mechanisms provides the capability for increased protection of individual system components and to more effectively control information flows between those components. Isolating system components provides enhanced protection that limits the potential harm from hostile cyber-attacks and errors. The degree of isolation varies depending upon the mechanisms chosen. Boundary protection mechanisms include routers, gateways, and firewalls that separate system components into physically separate networks or subnetworks; cross-domain devices that separate subnetworks; virtualization techniques; and the encryption of information flows among system components using distinct encryption keys.

boundary protection mechanisms are employed to isolate {{ insert: param, sc-07.21_odp.01 }} supporting {{ insert: param, sc-07.21_odp.02 }}.

System and communications protection policy

procedures addressing boundary protection

system design documentation

system hardware and software

enterprise architecture documentation

system architecture

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with boundary protection responsibilities

Mechanisms supporting and/or implementing the capability to separate system components supporting organizational missions and/or business functions

### SC-12 (1): Availability

Maintain availability of information in the event of the loss of cryptographic keys by users.

Escrowing of encryption keys is a common practice for ensuring availability in the event of key loss. A forgotten passphrase is an example of losing a cryptographic key.

information availability is maintained in the event of the loss of cryptographic keys by users.

System and communications protection policy

procedures addressing cryptographic key establishment, management, and recovery

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with responsibilities for cryptographic key establishment or management

Mechanisms supporting and/or implementing cryptographic key establishment and management

### SC-24: Fail in Known State

Fail to a {{ insert: param, sc-24_odp.02 }} for the following failures on the indicated components while preserving {{ insert: param, sc-24_odp.03 }} in failure: {{ insert: param, sc-24_odp.01 }}.

Failure in a known state addresses security concerns in accordance with the mission and business needs of organizations. Failure in a known state prevents the loss of confidentiality, integrity, or availability of information in the event of failures of organizational systems or system components. Failure in a known safe state helps to prevent systems from failing to a state that may cause injury to individuals or destruction to property. Preserving system state information facilitates system restart and return to the operational mode with less disruption of mission and business processes.

 {{ insert: param, sc-24_odp.01 }} fail to a {{ insert: param, sc-24_odp.02 }} while preserving {{ insert: param, sc-24_odp.03 }} in failure.

System and communications protection policy

procedures addressing system failure to known state

system design documentation

system configuration settings and associated documentation

list of failures requiring system to fail in a known state

state information to be preserved in system failure

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

Mechanisms supporting and/or implementing the fail in known state capability

mechanisms preserving system state information in the event of a system failure


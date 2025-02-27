/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package constants

// Common
const True = "true"
const False = "false"

// Cluster
const DefaultNodeAttributeHostNameKey = "si.io/hostname"
const DefaultNodeAttributeRackNameKey = "si.io/rackname"
const DefaultNodeInstanceTypeNodeLabelKey = "node.kubernetes.io/instance-type"
const DefaultRackName = "/rack-default"

// Application
const LabelApp = "app"
const LabelApplicationID = "applicationId"
const AnnotationApplicationID = "yunikorn.apache.org/app-id"
const LabelQueueName = "queue"
const RootQueue = "root"
const AnnotationQueueName = "yunikorn.apache.org/queue"
const AnnotationParentQueue = "yunikorn.apache.org/parentqueue"
const LabelDisableStateAware = "disableStateAware"
const ApplicationDefaultQueue = "root.sandbox"
const DefaultPartition = "default"
const AppTagNamespace = "namespace"
const AppTagNamespaceParentQueue = "namespace.parentqueue"
const AppTagImagePullSecrets = "imagePullSecrets"
const DefaultAppNamespace = "default"
const DefaultUserLabel = "yunikorn.apache.org/username"
const DefaultUser = "nobody"

// Spark
const SparkLabelAppID = "spark-app-selector"
const SparkLabelRole = "spark-role"
const SparkLabelRoleDriver = "driver"

// Configuration
const ConfigMapName = "yunikorn-configs"
const DefaultConfigMapName = "yunikorn-defaults"
const SchedulerName = "yunikorn"

// OwnerReferences
const DaemonSetType = "DaemonSet"

// Application crd
const AppManagerHandlerName = "yunikorn-app"

// Gang scheduling
const PlaceholderContainerImage = "registry.k8s.io/pause:3.7"
const PlaceholderContainerName = "pause"
const PlaceholderPodRestartPolicy = "Never"
const LabelPlaceholderFlag = "placeholder"
const AnnotationPlaceholderFlag = "yunikorn.apache.org/placeholder"
const AnnotationTaskGroupName = "yunikorn.apache.org/task-group-name"
const AnnotationTaskGroups = "yunikorn.apache.org/task-groups"
const AnnotationSchedulingPolicyParam = "yunikorn.apache.org/schedulingPolicyParameters"
const SchedulingPolicyTimeoutParam = "placeholderTimeoutInSeconds"
const SchedulingPolicyParamDelimiter = " "
const SchedulingPolicyStyleParam = "gangSchedulingStyle"
const SchedulingPolicyStyleParamDefault = "Soft"

var SchedulingPolicyStyleParamValues = map[string]string{"Hard": "Hard", "Soft": "Soft"}

const ApplicationInsufficientResourcesFailure = "ResourceReservationTimeout"
const ApplicationRejectedFailure = "ApplicationRejected"

// namespace.max.* (Retaining for backwards compatibility. Need to be removed in next major release)
const CPUQuota = "yunikorn.apache.org/namespace.max.cpu"
const MemQuota = "yunikorn.apache.org/namespace.max.memory"

// NamespaceQuota Namespace Quota
const NamespaceQuota = "yunikorn.apache.org/namespace.quota"

// NamespaceGuaranteed Namespace Guaranteed
const NamespaceGuaranteed = "yunikorn.apache.org/namespace.guaranteed"

// AnnotationAllowPreemption set on PriorityClass, opt out of preemption for pods with this priority class
const AnnotationAllowPreemption = "yunikorn.apache.org/allow-preemption"

// AnnotationIgnoreApplication set on Pod prevents by admission controller, prevents YuniKorn from honoring application ID
const AnnotationIgnoreApplication = "yunikorn.apache.org/ignore-application"

// AnnotationGenerateAppID adds application ID to workloads in the namespace even if not set in the admission config.
// Overrides the regexp behaviour if set, checked before the regexp is evaluated.
// true: add an application ID label
// false: do not add an application ID
const AnnotationGenerateAppID = "yunikorn.apache.org/namespace.generateAppId"

// AnnotationEnableYuniKorn sets the scheduler name to YuniKorn for workloads in the namespace even if not set in the admission config.
// Overrides the regexp behaviour if set, checked before the regexp is evaluated.
// true: set the scheduler name to YuniKorn
// false: do not do anything
const AnnotationEnableYuniKorn = "yunikorn.apache.org/namespace.enableYuniKorn"

// Admission Controller pod label update constants
const AutoGenAppPrefix = "yunikorn"
const AutoGenAppSuffix = "autogen"

// Compression Algorithms for schedulerConfig
const GzipSuffix = "gz"
